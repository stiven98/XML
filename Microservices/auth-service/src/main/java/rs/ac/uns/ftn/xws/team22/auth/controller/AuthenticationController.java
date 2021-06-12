package rs.ac.uns.ftn.xws.team22.auth.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;
import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationRequestDTO;
import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationResponseDTO;
import rs.ac.uns.ftn.xws.team22.auth.service.impl.AuthenticationService;

import javax.servlet.http.HttpServletRequest;

@RestController
@RequestMapping(value = "/auth", produces = MediaType.APPLICATION_JSON_VALUE)
public class AuthenticationController {

    @Autowired
    AuthenticationService authService;

    private static final Logger log = LoggerFactory.getLogger(AuthenticationController.class);

    @PostMapping
    public ResponseEntity<AuthenticationResponseDTO> createAuthenticationToken(@RequestBody AuthenticationRequestDTO dto, HttpServletRequest request) {
        try{
            AuthenticationResponseDTO dto1 = authService.login(dto);
            log.info("Successfully logged in "+ dto.getUsername());
            return new ResponseEntity<>(dto1, HttpStatus.OK);

        } catch (Exception e) {
            log.warn("Failed to log in " + dto.getUsername() + ", from: " + request.getHeader("Origin"));
            return new ResponseEntity<>(null, HttpStatus.OK);
        }
    }

    @GetMapping("/has-admin-role")
    @PreAuthorize("hasAuthority('ROLE_ADMIN')")
    public ResponseEntity<?> hasAdminRole(){
        try
        {
            log.info("Successfully got role  'ROLE_ADMIN'");
            return new ResponseEntity<>("", HttpStatus.OK);
        }
        catch (Exception e)
        {
            log.info("Failed to get role  'ROLE_ADMIN'");
            return new ResponseEntity<>("", HttpStatus.FORBIDDEN);
        }
    }

    @GetMapping("/has-system-user-role")
    @PreAuthorize("hasAuthority('ROLE_SYSTEM_USER')")
    public ResponseEntity<?> hasSystemUserRole(){
        try
        {
            log.info("Successfully got role  'ROLE_ADMIN'");
            return new ResponseEntity<>("", HttpStatus.OK);
        }
        catch (Exception e)
        {
            log.info("Failed to get role  'ROLE_ADMIN'");
            return new ResponseEntity<>("", HttpStatus.FORBIDDEN);
        }
    }

    @GetMapping("/has-agent-role")
    @PreAuthorize("hasAuthority('ROLE_AGENT')")
    public ResponseEntity<?> hasAgentRole(){
        try
        {
            log.info("Successfully got role  'ROLE_AGENT'");
            return new ResponseEntity<>("", HttpStatus.OK);
        }
        catch (Exception e)
        {
            log.info("Failed to get role  'ROLE_AGENT'");
            return new ResponseEntity<>("", HttpStatus.FORBIDDEN);
        }
    }

    @GetMapping("/is-authenticated")
    @PreAuthorize("isAuthenticated()")
    public ResponseEntity<?> isAuthenticated(){
        try
        {
            log.info("User is authenticated");
            return new ResponseEntity<>("", HttpStatus.OK);
        }
        catch (Exception e)
        {
            log.info("Failed to get role  'ROLE_AGENT'");
            return new ResponseEntity<>("", HttpStatus.FORBIDDEN);
        }
    }

}
