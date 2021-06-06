package rs.ac.uns.ftn.xws.team22.auth.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Service;
import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationRequestDTO;
import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationResponseDTO;
import rs.ac.uns.ftn.xws.team22.auth.model.AuthenticationData;
import rs.ac.uns.ftn.xws.team22.auth.security.TokenUtils;
import rs.ac.uns.ftn.xws.team22.auth.service.IAuthenticationService;

import java.util.Collection;

@Service
public class AuthenticationService implements IAuthenticationService {

    @Autowired
    private AuthenticationManager authenticationManager;

    @Autowired
    private TokenUtils tokenUtils;

    @Override
    public AuthenticationResponseDTO login(AuthenticationRequestDTO dto) {
        Authentication authentication = authenticationManager
                .authenticate(new UsernamePasswordAuthenticationToken(dto.getEmail(), dto.getPassword()));
        SecurityContextHolder.getContext().setAuthentication(authentication);
        AuthenticationData data = (AuthenticationData) authentication.getPrincipal();
        Collection<GrantedAuthority> authorities = (Collection<GrantedAuthority>) authentication.getAuthorities();
        String jwt = tokenUtils.generateToken(data.getUsername(), authorities);
        int expiresIn = tokenUtils.getExpiredIn();
        AuthenticationResponseDTO responseDTO = new AuthenticationResponseDTO(data.getUsername(), jwt, expiresIn);
        return responseDTO;
    }
}
