import React from "react";
import ProductsApi from "../api/ProductsApi";
import ImagesApi from "../api/ImagesApi";
import {withRouter} from "react-router-dom";
import ProductDetailModal from "./ProductDetailModal";
import ProductModelValidation from "../model/ProductModelValidation";
import OrdersApi from "../api/OrdersApi";

class ProductPage extends React.Component {

    constructor(props) {
        super(props);
        this.state = {product: null, showModal: false, quantity: 0, validQuantity: 'no-validate'}
    }

    componentDidMount = () => {
        if (localStorage.getItem('id') == null) {
            this.props.history.push('/forbidden');
            return;
        }

        const {id} = this.props.match.params;
        console.log(id);
        ProductsApi.get('get/' + id).then(response => {
            console.log(response);
            this.setState({product: response.data});
        }).catch(_ => {
            this.props.history.push('/home');
        });

    }

    deleteProduct = (id) => {
        return () => {
            ProductsApi.delete('delete/' + id).then((response => {
                console.log(response);
                this.props.history.push('/home');
            }))
        }

    }

    onChangeQuantity = (event) => {
        this.setState({quantity: event.target.value, validQuantity: 'no-validate'});

    }

    buyProduct = (event) => {
        event.preventDefault();

        if (this.state.quantity !== '') {
            this.setState({validQuantity: 'is-valid'});

            console.log("Salji req");

            const data = {
                quantity: parseInt(this.state.quantity),
                productID: this.state.product.id,
                userID: localStorage.getItem('id')

            };

            OrdersApi.post('create', data).then(response => {
                this.props.history.push('/myOrders')
            }).catch(err => {
               alert(err);
            });

        } else {
            this.setState({validQuantity: 'is-invalid'});
        }
    }

    renderButtons = () => {
        if (this.state.product.userID === localStorage.getItem('id')) {
            return (
                <div className={`pt-5 mt-5 d-flex justify-content-center`}>
                    <button className={`btn btn-success w-25 mr-2`} onClick={this.openModal}>Izmeni</button>
                    <button className={`btn btn-danger w-25 ml-2`}
                            onClick={this.deleteProduct(this.state.product.id)}>Obrisi
                    </button>
                </div>

            );
        } else {
            return (
                <form className="form-control-feedback mt-5">
                    <div className={`row`}>
                        <div className={`col-5 ml-3`}>
                            <label htmlFor={`quantity`} className={`text-dark`}>Kolicina:</label>
                            <input id={`quantity`} type={`number`} className={`form-control ` + this.state.validQuantity} value={this.state.quantity} onChange={this.onChangeQuantity}/>
                            <div className="invalid-feedback">
                                Incorrect value!
                            </div>
                        </div>
                        <div className={`col-6 mt-2`} >
                            <button className={`btn btn-success ml-3 mt-4 w-75`} onClick={this.buyProduct}>Kupi</button>
                        </div>
                    </div>

                </form>
            )


        }
    }

    updateProduct = async (product) => {
        await ProductsApi.put('/update', product).then((response) => {
            if (response.status === 200) {
                this.closeModal();
                this.componentDidMount();
                alert('Success!');
            }
        }).catch(err => {
            console.log(err);
            alert("Error");
        });
    }

    openModal = () => {
        this.setState({showModal: true});
    }

    closeModal = () => {
        this.setState({showModal: false});
        console.log(this.state.product);
    }

    render() {
        if (this.state.product != null) {
            return (
                <div>
                    <div className={`row mt-5 border border-dark rounded`}>
                        <div className={`col-6 mt-5 mb-5`}>
                            <div>
                                <img alt={`product`} src={ImagesApi + this.state.product.picturePath}
                                     style={{width: 500, height: 500}}/>
                            </div>
                        </div>

                        <div className={`col-6 my-auto`}>
                            <div className={`d-flex justify-content-center`}>
                                <h2 className={`text-dark`}>{this.state.product.name}</h2>
                            </div>

                            <div className={`d-flex justify-content-center`}>
                                <label className={`text-success`}>Cena: {this.state.product.price}</label>
                            </div>
                            <div className={`d-flex justify-content-center`}>
                                <label className={`text-success`}>Dostupna
                                    kolicina: {this.state.product.quantity}</label>
                            </div>
                            {this.renderButtons()}
                        </div>

                        <ProductDetailModal
                            showModal={this.state.showModal}
                            productModel={this.state.product}
                            productModelValidation={new ProductModelValidation()}
                            closeModal={this.closeModal}
                            saveButton={this.updateProduct}
                        />

                    </div>


                </div>
            );
        } else {
            return <div></div>;
        }
    }
}

export default withRouter(ProductPage);
