package rs.ac.uns.ftn.xws.team22.auth.dto;

import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

import java.util.UUID;

@Getter
@Setter
@NoArgsConstructor
public class BlockUserDTO {
    public UUID userId;
}
