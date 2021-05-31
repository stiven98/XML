package rs.ac.uns.ftn.xws.team22.auth.service;

import rs.ac.uns.ftn.xws.team22.auth.model.LoginDetails;

import java.util.List;
import java.util.UUID;

public interface ILoginDetailsService {
    List<LoginDetails> findAll();
    LoginDetails findById(UUID id);
    LoginDetails findByEmail(String email);
}
