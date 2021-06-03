package rs.ac.uns.ftn.xws.team22.auth.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
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
}
