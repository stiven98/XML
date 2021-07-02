
import React from "react";

class Forbidden extends React.Component {

    onClickHomeButton = () => {
        this.props.history.push('/home');
    }

    onClickLoginButton = () => {
        this.props.history.push('/login');
    }

    renderButtons = () => {
        if (localStorage.getItem('id') == null) {
            return (<button className={`btn btn-dark`} onClick={this.onClickLoginButton}>Go on Login page</button>);
        } else {
            return (<button className={`btn btn-dark`} onClick={this.onClickHomeButton}>Go on Home page</button>)
        }
    }

    render() {
        return (
            <div className={`m-5`}>
                <h2 className={`text-dark`}>Forbidden error!</h2>
                {this.renderButtons()}
            </div>
        );
    }
}

export default Forbidden;
