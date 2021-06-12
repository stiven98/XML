package rs.ac.uns.ftn.xws.team22.auth.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.*;
import rs.ac.uns.ftn.xws.team22.auth.dto.CreateUserCredentialsDTO;
import rs.ac.uns.ftn.xws.team22.auth.email.EmailSender;
import rs.ac.uns.ftn.xws.team22.auth.model.AuthenticationData;
import rs.ac.uns.ftn.xws.team22.auth.repository.LoginDetailsRepository;
import rs.ac.uns.ftn.xws.team22.auth.service.impl.AuthenticationDataService;

import javax.servlet.http.HttpServletRequest;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping(value = "/api/login-details", produces = MediaType.APPLICATION_JSON_VALUE)
public class LoginDetailsController {

    @Autowired
    AuthenticationDataService loginDetailsService;

    @Autowired
    EmailSender emailSender;

    @Autowired
    LoginDetailsRepository loginDetailsRepository;

    @Autowired
    PasswordEncoder passwordEncoder;

    private static final Logger log = LoggerFactory.getLogger(LoginDetailsController.class);

    @GetMapping
    public ResponseEntity<?> findAll() {
        Map<String, String> result = new HashMap<>();
        try
        {
            List<AuthenticationData> authenticationDataList = loginDetailsService.findAll();
            log.info("Successfully got all authenticated data");
            return new ResponseEntity<>(authenticationDataList, HttpStatus.OK);
        }
        catch (Exception e)
        {
            log.info("Failed to get all authenticated data");
            return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
        }
    }

    @PostMapping
    public ResponseEntity<?> createUserCredentials(@RequestBody CreateUserCredentialsDTO dto, HttpServletRequest request) {
        AuthenticationData authenticationData = new AuthenticationData();
        authenticationData.setActive(false);
        authenticationData.setEmail(dto.getEmail());
        authenticationData.setPassword(passwordEncoder.encode(dto.getPassword()));
        authenticationData.setUsername(dto.getUsername());
        AuthenticationData existEmail = loginDetailsRepository.findByEmail(dto.getEmail());
        if(existEmail != null){
            log.warn("Failed to create user credentials with email: " + dto.getEmail() + ", from: " + request.getHeader("Origin"));
            return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
        }
        AuthenticationData existUsername = loginDetailsRepository.findByUsername(dto.getUsername());
        if(existUsername != null){
            log.warn("Failed to create user credentials with username: " + dto.getUsername() + ", from: " + request.getHeader("Origin"));
            return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
        }
        AuthenticationData result = loginDetailsRepository.save(authenticationData);
        if(result == null){
            log.warn("Failed to create user credentials with username: " + dto.getUsername() + ", from: " + request.getHeader("Origin"));
            return new ResponseEntity<>(null, HttpStatus.BAD_REQUEST);
        }
        try {
            emailSender.sendActivationEmail(dto.getEmail(), result.getId().toString());
        }catch (Exception e){
            log.warn("Failed to send email");
        }
        log.info("Successfully created authentication data for: "+ dto.getUsername());
        return new ResponseEntity<>(null, HttpStatus.OK);
    }
}
