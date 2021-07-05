import React from "react";
import LoginInfoModel from "../model/LoginInfoModel";
import LoginInfoModelValidation from "../model/LoginInfoModelValidation";
import backendApi from "../api/UsersApi";


class Login extends React.Component {

    constructor(props) {
        super(props);
        this.state = {loginInfoModel: new LoginInfoModel(), loginInfoModelValidation: new LoginInfoModelValidation()};
    }

    componentDidMount() {
        console.log(localStorage.getItem("id"));
        if (localStorage.getItem("id") !== null) {
            this.props.history.push('/forbidden');
        }
    }

    onChangeUsernameInput = (event) => {
        this.setState({loginInfoModel : {...(this.state.loginInfoModel), username: event.target.value} });
        this.setState({loginInfoModelValidation: new LoginInfoModelValidation()});
    }

    onChangePasswordInput = (event) => {
        this.setState({loginInfoModel : {...(this.state.loginInfoModel), password: event.target.value} });
        this.setState({loginInfoModelValidation: new LoginInfoModelValidation()});
    }

    loginUser =  async (event) => {
        event.preventDefault();
        console.log(this.state.loginInfoModel);

        if (await this.isFormValid()) {
            await backendApi.post('/login', this.state.loginInfoModel).then((response) => {
                localStorage.setItem("id", response.data.id);
                localStorage.setItem("username", response.data.username);
                this.props.history.push("/home")
            }).catch((error) => {
                alert("Wrong data!");
            });

        }
    }

    isFormValid = async () => {
        const validEmail = await this.isValidEmail(this.state.loginInfoModel.username);
        const validPassword = await this.isValidPassword(this.state.loginInfoModel.password);
        return validEmail && validPassword;
    }


    isValidEmail = (email) => {
        if (email.length > 4) {
            this.setState({loginInfoModelValidation : {...this.state.loginInfoModelValidation, validUsername: 'is-valid'}});
            return true;
        } else {
            this.setState({loginInfoModelValidation : {...this.state.loginInfoModelValidation, validUsername: 'is-invalid'}});
            return false;
        }
    }

    isValidPassword = (password) => {
        if (password.length > 4) {
            this.setState({loginInfoModelValidation : {...this.state.loginInfoModelValidation, validPassword: 'is-valid'}});
            return true;
        } else {
            this.setState({loginInfoModelValidation : {...this.state.loginInfoModelValidation, validPassword: 'is-invalid'}});
            return false;
        }
    }

    goOnRegistration = () => {
        this.props.history.push('/registration');
    }

    render() {
        return (
            <div className={`container`}>
                <h1 className={`text text-dark font-weight-bold d-flex justify-content-center m-5`}>Login:</h1>

                <div className={`d-flex justify-content-center`}>
                    <form className="form-control-feedback w-50">
                        <div className="row d-flex justify-content-center mt-3">
                            <div className="col-8">
                                <div>
                                    <label htmlFor="email" className="text-dark">Email:</label>
                                    <input type="text" className={`form-control ` + this.state.loginInfoModelValidation.validUsername} required={true} id="email" value={this.state.loginInfoModel.username} onChange={this.onChangeUsernameInput}/>
                                    <div className="invalid-feedback">
                                        Incorrect username!
                                    </div>
                                </div>
                                <div>
                                    <label className="pt-3 text-dark" htmlFor="password">Password:</label>
                                    <input type="password" className={`form-control ` + this.state.loginInfoModelValidation.validPassword} required={true} id="password" value={this.state.loginInfoModel.password} onChange={this.onChangePasswordInput}/>
                                    <div className="invalid-feedback">
                                        Minimum 8 characters.
                                    </div>
                                </div>

                                <button  type="submit" className="btn btn-dark mt-3 w-100" onClick={this.loginUser}>Submit</button>
                                <button  type="submit" className="btn btn-dark mt-3 w-100" onClick={this.goOnRegistration}>Go on Registration</button>

                            </div>
                        </div>
                    </form>
                </div>
            </div>


        );
    }


}

export default Login;
