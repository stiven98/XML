package rs.ac.uns.ftn.xws.team22.auth.controller;


import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.MediaType;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import rs.ac.uns.ftn.xws.team22.auth.dto.BlockUserDTO;
import rs.ac.uns.ftn.xws.team22.auth.dto.CreateUserDTO;
import rs.ac.uns.ftn.xws.team22.auth.model.AuthenticationData;
import rs.ac.uns.ftn.xws.team22.auth.service.impl.AuthenticationDataService;

import java.util.UUID;


@RestController
@RequestMapping(value = "/api", produces = MediaType.APPLICATION_JSON_VALUE)
public class AuthenticationDataController {

    @Autowired
    AuthenticationDataService authenticationDataService;


    @PostMapping("/blockUser")
    public ResponseEntity<?> blockUser(@RequestBody BlockUserDTO blockUserDTO) {
        boolean isBlocked = authenticationDataService.blockUser(blockUserDTO.userId);
        return  new ResponseEntity<>(isBlocked, HttpStatus.OK);
    }

    @PostMapping("/createUser")
    public ResponseEntity<?> createUser(@RequestBody CreateUserDTO dto) {
        AuthenticationData authData = authenticationDataService.createUser(dto);
        return  new ResponseEntity<>(authData, HttpStatus.CREATED);
    }

}
