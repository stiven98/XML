package rs.ac.uns.ftn.xws.team22.auth.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.stereotype.Service;
import rs.ac.uns.ftn.xws.team22.auth.model.AuthenticationData;
import rs.ac.uns.ftn.xws.team22.auth.repository.LoginDetailsRepository;
import rs.ac.uns.ftn.xws.team22.auth.service.IAuthenticationDataService;

import java.util.List;
import java.util.UUID;

@Service
public class AuthenticationDataService implements IAuthenticationDataService, UserDetailsService {

    @Autowired
    private LoginDetailsRepository loginDetailsRepository;

    @Override
    public List<AuthenticationData> findAll() {
        return this.loginDetailsRepository.findAll();
    }

    @Override
    public AuthenticationData findById(UUID id) {
        return this.loginDetailsRepository.getById(id);
    }

    @Override
    public AuthenticationData findByEmail(String email) {
        return this.loginDetailsRepository.findByEmail(email);
    }

    @Override
    public UserDetails loadUserByUsername(String s) throws UsernameNotFoundException {
        return this.loginDetailsRepository.findByEmail(s);
    }
}
