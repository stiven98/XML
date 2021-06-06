package rs.ac.uns.ftn.xws.team22.auth.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class AuthenticationResponseDTO {
    public String username;
    public String token;
    public String role;
    public int expiresIn;
}
