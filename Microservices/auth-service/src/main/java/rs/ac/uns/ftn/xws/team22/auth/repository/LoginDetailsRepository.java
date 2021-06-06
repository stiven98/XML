package rs.ac.uns.ftn.xws.team22.auth.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import rs.ac.uns.ftn.xws.team22.auth.model.AuthenticationData;

import java.util.UUID;

public interface LoginDetailsRepository extends JpaRepository<AuthenticationData, UUID> {

    AuthenticationData findByUsername(String username);

    @Query("select ad.username from AuthenticationData ad where ad.username = ?1")
    String isValidUsername(String username);

    @Query("select ad.password from AuthenticationData ad where ad.username = ?1")
    String isValidPassword(String username);

}
