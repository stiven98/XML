package rs.ac.uns.ftn.xws.team22.auth.service.impl;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import rs.ac.uns.ftn.xws.team22.auth.controller.AuthenticationController;
import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationRequestDTO;
import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationResponseDTO;
import rs.ac.uns.ftn.xws.team22.auth.dto.ResetPasswordDTO;
import rs.ac.uns.ftn.xws.team22.auth.email.EmailSender;
import rs.ac.uns.ftn.xws.team22.auth.model.AuthenticationData;
import rs.ac.uns.ftn.xws.team22.auth.model.ResetPasswordRequest;
import rs.ac.uns.ftn.xws.team22.auth.repository.LoginDetailsRepository;
import rs.ac.uns.ftn.xws.team22.auth.repository.ResetPasswordRequestRepository;
import rs.ac.uns.ftn.xws.team22.auth.security.TokenUtils;
import rs.ac.uns.ftn.xws.team22.auth.service.IAuthenticationService;

import java.util.Calendar;
import java.util.Collection;
import java.util.Date;
import java.util.UUID;

@Service
public class AuthenticationService implements IAuthenticationService {

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private TokenUtils tokenUtils;

    @Autowired
    private LoginDetailsRepository repository;

    @Autowired
    private EmailSender emailSender;

    @Autowired
    PasswordEncoder passwordEncoder;

    @Autowired
    private ResetPasswordRequestRepository passwordRequestRepository;

    private static final Logger log = LoggerFactory.getLogger(AuthenticationService.class);

    @Override
    public AuthenticationResponseDTO login(AuthenticationRequestDTO dto) {
        Authentication authentication = authenticationManager
                .authenticate(new UsernamePasswordAuthenticationToken(dto.getUsername(), dto.getPassword()));
        SecurityContextHolder.getContext().setAuthentication(authentication);
        AuthenticationData data = (AuthenticationData) authentication.getPrincipal();
        Collection<GrantedAuthority> authorities = (Collection<GrantedAuthority>) authentication.getAuthorities();
        String jwt = tokenUtils.generateToken(data.getUsername(), authorities);
        int expiresIn = tokenUtils.getExpiredIn();
        String role = null;
        for (GrantedAuthority grantedAuthority : authorities) {
            if (grantedAuthority.getAuthority().contains("ROLE")) {
                role = grantedAuthority.getAuthority();
                break;
            }
        }
        AuthenticationResponseDTO responseDTO = new AuthenticationResponseDTO(data.getId(), data.getUsername(), jwt, role, expiresIn);
        return responseDTO;
    }

    @Override
    public boolean sendResetPasswordRequest(String email) {
        AuthenticationData user= repository.findByEmail(email);
        if(user==null){
            log.error("Send reset password error,user do not exist:" +email );
            return false;
        }
        ResetPasswordRequest request=new ResetPasswordRequest();
        request.setEmail(email);
        request.setUsed(false);
        Date date=new Date();
        Calendar cal = Calendar.getInstance();
        cal.setTime(date);
        cal.add(Calendar.DATE, 1);
        request.setValidTo(cal.getTime());
        ResetPasswordRequest request1=passwordRequestRepository.save(request);
        try {
            emailSender.sendForgotPasswordEmail(request1.getId().toString(),email);
        } catch(Exception e) {
            log.error(e.getMessage());
            return false;
        }
        return true;
    }

    @Override
    public boolean resetPassword(ResetPasswordDTO dto) {
        if(!dto.getPassword().equals(dto.getPassword2())){
            return false;
        }
        ResetPasswordRequest request=checkRequest(dto.getRequestId());
        if(request!=null){
            AuthenticationData user= repository.findByEmail(request.getEmail());
            if(user!=null){
                user.setPassword(passwordEncoder.encode(dto.getPassword()));
                repository.save(user);
                request.setUsed(true);
                passwordRequestRepository.save(request);
                try {
                    emailSender.sendResetPasswordEmail(request.getEmail());
                }catch (Exception e){
                    log.warn("Failed to send email");
                    return false;
                }
                log.info("Password successfully changed user :" +user.getUsername());
                return true;
            }
        }
        log.warn("Failed to reset password invalid request:"+ dto.getRequestId() );
        return false;
    }

    @Override
    public ResetPasswordRequest checkRequest(UUID id) {
        ResetPasswordRequest request= passwordRequestRepository.findById(id).orElse(null);
        if(request!=null){
            if(!request.isUsed() && !request.getValidTo().before(new Date())) {
                return  request;
            }
        }
        return null;
    }
}
