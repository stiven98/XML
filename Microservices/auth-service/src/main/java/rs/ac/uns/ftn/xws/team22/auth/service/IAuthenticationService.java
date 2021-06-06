package rs.ac.uns.ftn.xws.team22.auth.service;

import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationRequestDTO;
import rs.ac.uns.ftn.xws.team22.auth.dto.AuthenticationResponseDTO;

public interface IAuthenticationService {
    AuthenticationResponseDTO login(AuthenticationRequestDTO dto);
}
