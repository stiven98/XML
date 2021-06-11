package rs.ac.uns.ftn.xws.team22.auth.controller;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import rs.ac.uns.ftn.xws.team22.auth.model.AuthenticationData;
import rs.ac.uns.ftn.xws.team22.auth.service.impl.AuthenticationDataService;

import java.util.HashMap;
import java.util.List;
import java.util.Map;

@RestController
@RequestMapping(value = "/api/login-details", produces = MediaType.APPLICATION_JSON_VALUE)
public class LoginDetailsController {

    @Autowired
    AuthenticationDataService loginDetailsService;

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


}
