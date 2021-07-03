import React from "react";
import OrdersApi from "../api/OrdersApi";
import { withRouter } from "react-router-dom";

class MyOrders extends React.Component {

    constructor(props) {
        super(props);
        this.state = {orders: []}
    }

    componentDidMount = () => {
        if (localStorage.getItem('id') == null) {
            this.props.history.push('/forbidden');
            return;
        }

        OrdersApi.get('get/' + localStorage.getItem('id')).then((response) => {
           this.setState({orders: response.data});
        });
    }

    renderRows = () => {
        return this.state.orders.map((item, index) => {
            return (
                    <tr>
                        <th scope="row">{index + 1}</th>
                        <td>{item.Product.name}</td>
                        <td>{item.quantity}</td>

                    </tr>
                );



        })
    }

    render() {
        if (this.state.orders.length === 0) {
            return (
                <div className={`m-5`}>
                    <h2 className={`text-dark`}>Nema kupovina!</h2>
                </div>
            );
        } else {
            return (
                <div>
                    <div className={`m-5`}>
                        <h2 className={`text-dark`}>Kupovine:</h2>
                    </div>

                    <table className="table table-striped table-dark">
                        <thead>
                            <th scope="col">#</th>
                            <th scope="col">Naziv proizvoda</th>
                            <th scope="col">Kolicina</th>
                        </thead>
                        <tbody>
                        {this.renderRows()}
                        </tbody>

                    </table>




                </div>

            );
        }

    }
}

export default withRouter(MyOrders);
