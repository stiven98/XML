import React from "react";
import ProductsApi from "../api/ProductsApi";
import ImagesApi from "../api/ImagesApi";
import { withRouter } from "react-router-dom";
import ProductDetailModal from "./ProductDetailModal";
import ProductModel from "../model/ProductModel";
import ProductModelValidation from "../model/ProductModelValidation";

class ProductPage extends React.Component {

    constructor(props) {
        super(props);
        this.state = {product: null, showModal: false}
    }



    componentDidMount = () => {
        if (localStorage.getItem('id') == null) {
            this.props.history.push('/forbidden');
            return;
        }

        const { id } = this.props.match.params;
        console.log(id);
        ProductsApi.get('get/' + id).then(response => {
            console.log(response);
            this.setState({product: response.data});
        }).catch(_ => {
            this.props.history.push('/home');
        });

    }

    deleteProduct =  (id) => {
        return () => {
            ProductsApi.delete('delete/' + id).then((response => {
                console.log(response);
                this.props.history.push('/home');
            }))
        }

    }

    renderButtons = () => {
        if (this.state.product.userID === localStorage.getItem('id')) {
            return (
                <div className={`pt-5 mt-5 d-flex justify-content-center`}>
                    <button className={`btn btn-success w-25 mr-2`} onClick={this.openModal}>Izmeni</button>
                    <button className={`btn btn-danger w-25 ml-2`} onClick={this.deleteProduct(this.state.product.id)}>Obrisi</button>
                </div>

            );
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
                            <div >
                                <img src={ImagesApi + this.state.product.picturePath}/>
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
                                <label className={`text-success`}>Dostupna kolicina: {this.state.product.quantity}</label>
                            </div>
                            {this.renderButtons()}
                        </div>

                        <ProductDetailModal
                            showModal={this.state.showModal}
                            productModel= {this.state.product}
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
