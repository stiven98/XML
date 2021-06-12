package rs.ac.uns.ftn.xws.team22.auth.email;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.stereotype.Component;

@Component
public class EmailSender {


    @Autowired
    private JavaMailSender emailSender;

    public void sendForgotPasswordEmail(String id,String mail) {
        SimpleMailMessage message = new SimpleMailMessage();
        String body = String.format("Please click on the link to reset your password: http://localhost:4200/reset/"+id);
        message.setTo(mail);
        message.setSubject("Forgot password");
        message.setText(body);
        emailSender.send(message);
    }

    public void sendResetPasswordEmail(String mail){
        SimpleMailMessage message = new SimpleMailMessage();
        String body = String.format("Your password is successfully changed, if you didn't change it please contact administrator. ");
        message.setTo(mail);
        message.setSubject("Changed password");
        message.setText(body);
        emailSender.send(message);
    }

    public void sendActivationEmail(String mail,String id){
        SimpleMailMessage message = new SimpleMailMessage();
        String body = String.format("Please click on the link to activate your account: http://localhost:4200/activate/"+id );
        message.setTo(mail);
        message.setSubject("Activate account");
        message.setText(body);
        emailSender.send(message);
    }
}
