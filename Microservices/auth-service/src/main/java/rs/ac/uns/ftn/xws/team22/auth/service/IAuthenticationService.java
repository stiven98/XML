package rs.ac.uns.ftn.xws.team22.auth.service;

import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationRequestDTO;
import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationResponseDTO;
import rs.ac.uns.ftn.xws.team22.auth.dto.ResetPasswordDTO;
import rs.ac.uns.ftn.xws.team22.auth.model.ResetPasswordRequest;

import java.util.UUID;

public interface IAuthenticationService {
    AuthenticationResponseDTO login(AuthenticationRequestDTO dto);

    boolean activateAccount(UUID id);

    ResetPasswordRequest checkRequest(UUID id);


    boolean sendResetPasswordRequest(String email);

    boolean resetPassword(ResetPasswordDTO dto);
}
