import {BrowserRouter, Route} from "react-router-dom";
import Forbidden from "./Forbidden";
import Home from "./Home";
import Login from "./Login";
import Registration from "./Registration";
import Header from "./Header";
import MyProducts from "./MyProducts";
import ProductPage from "./ProductPage";
import MyOrders from "./MyOrders";
import Campaign from "./Campaign";

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
                    <Route path={`/myOrders`} exact={true} component={MyOrders} />
                    <Route path={'/login'} exact={true} component={Login} />
                    <Route path={'/forbidden'} exact={true} component={Forbidden} />
                    <Route path={'/registration'} exact={true} component={Registration} />
                    <Route path={'/campaign'} exact={true} component={Campaign} />
                    <Route path={'/product/:id'} exact={true} component={ProductPage} />
                </div>
            </BrowserRouter>
        </div>
    );
}

export default App;
