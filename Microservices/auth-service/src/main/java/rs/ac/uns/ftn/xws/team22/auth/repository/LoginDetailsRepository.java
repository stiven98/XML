package rs.ac.uns.ftn.xws.team22.auth.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import rs.ac.uns.ftn.xws.team22.auth.model.AuthenticationData;

import java.util.UUID;

public interface LoginDetailsRepository extends JpaRepository<AuthenticationData, UUID> {

    AuthenticationData findByEmail(String email);

}
