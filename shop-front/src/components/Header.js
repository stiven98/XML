import React from "react";

import {withRouter} from 'react-router-dom';

class Header extends React.Component {


    renderHeader = () => {
        console.log(localStorage.getItem('id'));
        if (localStorage.getItem('id') != null) {
            return (<Header /> );
        }
    }

    logout = () => {
        localStorage.removeItem('id');
        localStorage.removeItem('username');
        this.props.history.push("/login");
    }

    allProducts = () => {
        this.props.history.push("/home")
    }

    myProducts = () => {
        this.props.history.push("/myProducts")
    }

    render() {
        if (localStorage.getItem('id') != null) {
            return (
                <div>
                    <nav className="navbar navbar-expand-lg navbar-dark bg-dark text-dark">
                        <a className="navbar-brand" href="/">Shop</a>
                        <div className="container-fluid nav nav-pills" role="tablist">
                            <button className="btn bg-light ml-2 mr-2" onClick={this.allProducts}>Svi proizvodi</button>
                            <button className="btn bg-light ml-2 mr-auto" onClick={this.myProducts}>Moji proizvodi</button>
                            <button className="btn bg-light ml-2 mr-2" onClick={this.logout}>Odjavi se</button>
                        </div>
                    </nav>
                </div>
            );
        } else {
            return (<div></div>);
        }

    }


}

export default withRouter(Header);
