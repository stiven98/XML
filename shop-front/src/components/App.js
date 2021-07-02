import {BrowserRouter, Route} from "react-router-dom";
import Forbidden from "./Forbidden";
import Home from "./Home";
import Login from "./Login";
import Registration from "./Registration";
import Header from "./Header";
import MyProducts from "./MyProducts";
import ProductPage from "./ProductPage";

function App() {

    return (
        <div>
            <BrowserRouter>
                <div>
                    <Header/>
                </div>
                <div className={`container`}>
                    <Route path={`/`} exact={true} component={Home} />
                    <Route path={`/home`} exact={true} component={Home}/>
                    <Route path={`/myProducts`} exact={true} component={MyProducts} />
                    <Route path={'/login'} exact={true} component={Login} />
                    <Route path={'/forbidden'} exact={true} component={Forbidden} />
                    <Route path={'/registration'} exact={true} component={Registration} />
                    <Route path={'/product/:id'} exact={true} component={ProductPage} />
                </div>
            </BrowserRouter>
        </div>
    );
}

export default App;
