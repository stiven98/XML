package rs.ac.uns.ftn.xws.team22.auth.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.core.userdetails.UsernameNotFoundException;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.stereotype.Service;
import rs.ac.uns.ftn.xws.team22.auth.dto.CreateUserDTO;
import rs.ac.uns.ftn.xws.team22.auth.model.AuthenticationData;
import rs.ac.uns.ftn.xws.team22.auth.model.Role;
import rs.ac.uns.ftn.xws.team22.auth.repository.LoginDetailsRepository;
import rs.ac.uns.ftn.xws.team22.auth.repository.RoleRepository;
import rs.ac.uns.ftn.xws.team22.auth.service.IAuthenticationDataService;

import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

@Service
public class AuthenticationDataService implements IAuthenticationDataService, UserDetailsService {

    @Autowired
    private LoginDetailsRepository loginDetailsRepository;
    @Autowired
    private RoleRepository roleRepository;
    @Autowired
    private PasswordEncoder passwordEncoder;
    @Override
    public List<AuthenticationData> findAll() {
        return this.loginDetailsRepository.findAll();
    }

    @Override
    public AuthenticationData findById(UUID id) {
        return this.loginDetailsRepository.getById(id);
    }

    @Override
    public AuthenticationData findByUsername(String email) {
        return this.loginDetailsRepository.findByUsername(email);
    }

    @Override
    public boolean isValidLogin(String username, String password) {
        String validPassword = loginDetailsRepository.isValidPassword(username);
        if(loginDetailsRepository.isValidUsername(username) != null && passwordEncoder.matches(password, validPassword) ) {
            return true;
        }
        else {
            return false;
        }
    }

    @Override
    public UserDetails loadUserByUsername(String s) throws UsernameNotFoundException {
        return this.loginDetailsRepository.findByUsername(s);
    }
    @Override
    public boolean blockUser(UUID userId) {
        AuthenticationData userDetails = this.loginDetailsRepository.getById(userId);
        if ( userDetails == null) {
            return false;
        }
        userDetails.setActive(false);
        this.loginDetailsRepository.save(userDetails);
        return true;
    }

    @Override
    public AuthenticationData createUser(CreateUserDTO dto) {
        AuthenticationData data = new AuthenticationData();
        data.setUsername(dto.username);
        data.setPassword(dto.password);
        data.setActive(dto.isActive);
        data.setId(dto.id);
        List<Role> roles = new ArrayList<>();
        Role role = this.roleRepository.findRoleByRole(dto.getRole());
        roles.add(role);
        data.setRoles(roles);
        this.loginDetailsRepository.save(data);
        return data;
    }


}
