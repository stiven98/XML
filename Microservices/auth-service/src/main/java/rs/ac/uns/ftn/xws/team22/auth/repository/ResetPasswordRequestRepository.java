package rs.ac.uns.ftn.xws.team22.auth.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import rs.ac.uns.ftn.xws.team22.auth.model.ResetPasswordRequest;

import java.util.UUID;

public interface ResetPasswordRequestRepository extends JpaRepository<ResetPasswordRequest, UUID> {
}
