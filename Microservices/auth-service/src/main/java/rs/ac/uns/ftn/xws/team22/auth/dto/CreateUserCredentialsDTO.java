package rs.ac.uns.ftn.xws.team22.auth.dto;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class CreateUserCredentialsDTO {
    public String username;
    public String password;
    public String email;
}
