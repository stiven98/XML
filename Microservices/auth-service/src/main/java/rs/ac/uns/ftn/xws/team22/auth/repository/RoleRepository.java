package rs.ac.uns.ftn.xws.team22.auth.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import rs.ac.uns.ftn.xws.team22.auth.model.Role;

import java.util.UUID;

public interface RoleRepository extends JpaRepository<Role, UUID> {

    Role findRoleByRole(String role);
}
