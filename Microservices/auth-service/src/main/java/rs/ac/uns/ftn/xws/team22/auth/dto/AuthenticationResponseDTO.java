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
    public String email;
    public String token;
    public int expiresIn;
}
