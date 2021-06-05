package rs.ac.uns.ftn.xws.team22.auth.dto;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@NoArgsConstructor
public class AuthenticationRequestDTO {
    public String email;
    public String password;
}
