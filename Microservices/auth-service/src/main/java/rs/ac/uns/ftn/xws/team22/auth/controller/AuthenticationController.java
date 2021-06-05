package rs.ac.uns.ftn.xws.team22.auth.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.web.bind.annotation.*;
import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationRequestDTO;
import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationResponseDTO;
import rs.ac.uns.ftn.xws.team22.auth.service.impl.AuthenticationService;

@RestController
@RequestMapping(value = "/auth", produces = MediaType.APPLICATION_JSON_VALUE)
public class AuthenticationController {

    @Autowired
    AuthenticationService authService;

    @PostMapping
    public ResponseEntity<AuthenticationResponseDTO> createAuthenticationToken(@RequestBody AuthenticationRequestDTO dto) {
        return new ResponseEntity<>(authService.login(dto), HttpStatus.OK);
    }

    @GetMapping("/has-admin-role")
    @PreAuthorize("hasAuthority('ROLE_ADMIN')")
    public ResponseEntity<?> hasAdminRole(){
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    @GetMapping("/has-system-user-role")
    @PreAuthorize("hasAuthority('ROLE_SYSTEM_USER')")
    public ResponseEntity<?> hasSystemUserRole(){
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    @GetMapping("/has-agent-role")
    @PreAuthorize("hasAuthority('ROLE_AGENT')")
    public ResponseEntity<?> hasAgentRole(){
        return new ResponseEntity<>("", HttpStatus.OK);
    }

    @GetMapping("/is-authenticated")
    @PreAuthorize("isAuthenticated()")
    public ResponseEntity<?> isAuthenticated(){
        return new ResponseEntity<>("", HttpStatus.OK);
    }

}
