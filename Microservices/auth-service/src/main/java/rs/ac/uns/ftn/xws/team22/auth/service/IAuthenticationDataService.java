package rs.ac.uns.ftn.xws.team22.auth.service;

import rs.ac.uns.ftn.xws.team22.auth.model.AuthenticationData;

import java.util.List;
import java.util.UUID;

public interface IAuthenticationDataService {
    List<AuthenticationData> findAll();
    AuthenticationData findById(UUID id);
    AuthenticationData findByEmail(String email);
}
