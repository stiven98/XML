package rs.ac.uns.ftn.xws.team22.auth.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.*;
import rs.ac.uns.ftn.xws.team22.auth.model.AuthenticationData;
import rs.ac.uns.ftn.xws.team22.auth.service.impl.AuthenticationDataService;

import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.UUID;

@RestController
@RequestMapping(value = "/api/login-details", produces = MediaType.APPLICATION_JSON_VALUE)
public class LoginDetailsController {

    @Autowired
    AuthenticationDataService loginDetailsService;

    @Autowired
    PasswordEncoder passwordEncoder;

    @GetMapping
    public ResponseEntity<?> findAll() {
        Map<String, String> result = new HashMap<>();
        List<AuthenticationData> authenticationDataList = loginDetailsService.findAll();
        return new ResponseEntity<>(authenticationDataList, HttpStatus.OK);
    }

    @GetMapping("/isValidLogin/{username}/{password}")
    public ResponseEntity<?> isValidLogin(@PathVariable String username, @PathVariable String password) {
        boolean isValidLogin = loginDetailsService.isValidLogin(username, password);
        return new ResponseEntity<Boolean>(isValidLogin, HttpStatus.OK);
    }
}
