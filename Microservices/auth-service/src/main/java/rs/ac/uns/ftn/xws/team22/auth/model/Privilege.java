package rs.ac.uns.ftn.xws.team22.auth.model;
import com.fasterxml.jackson.annotation.JsonIgnore;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;
import org.hibernate.annotations.GenericGenerator;

import javax.persistence.*;
import java.util.Collection;
import java.util.UUID;

@Getter
@Setter
@NoArgsConstructor
@Entity
@Table(name = "PRIVILEGES")
public class Privilege{
    @Id
    @GeneratedValue(generator = "uuid2")
    @GenericGenerator(name = "uuid2", strategy = "uuid2")
    @Column(name = "privilege_id", nullable = false, unique = true)
    private UUID privilegeId;

    @Column(name="name")
    private String name;

    @ManyToMany(mappedBy = "privileges")
    @JsonIgnore
    private Collection<Role> roles;

}