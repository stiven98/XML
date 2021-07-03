import React from "react";
import ProductsApi from "../api/ProductsApi";
import ProductCard from "./ProductCard";
import {Button, Modal} from "react-bootstrap";
import ProductModel from "../model/ProductModel";
import ProductModelValidation from "../model/ProductModelValidation";
import ImagesApi from "../api/ImagesApi";
import axios from "axios";
import usersApi from "../api/UsersApi";
import UserModel from "../model/UserModel";
import UserModelValidation from "../model/UserModelValidation";
import ProductDetailModal from "./ProductDetailModal";

class MyProducts extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            products: [],
            showModal: false,
            productModel: new ProductModel(),
            productModelValidation: new ProductModelValidation(),
        };
    }

    componentDidMount = async () => {
        if (localStorage.getItem('id') === null) {
            this.props.history.push('/forbidden');
        }

        await ProductsApi.get('/user/' + localStorage.getItem('id')).then(response => {
            this.setState({products: response.data})
            console.log(this.state.products);
        })


    }

    createProduct = async (product) => {
        await ProductsApi.post('/create', product).then((response) => {
                if (response.status === 201) {
                    this.closeModal();
                    this.componentDidMount();
                    alert('Success!');
                }
            }).catch(err => {
                console.log(err);
                alert("Error");
            });
    }

    renderProductsCards = () => {
        return this.state.products.map(product => {
            if (product.deleted === false) {
                return (<ProductCard key={product.id} product={product}/>);
            }
        })
    }

    openModal = () => {
        this.setState({showModal: true,  productModel: new ProductModel()});
    }

    closeModal = () => {
        this.setState({showModal: false});
    }

    render() {
        return (
            <div>
                <div className={`row mt-4`}>
                    <div className={`col d-flex justify-content-start`}>
                        <h2 className={`text-dark font-weight-bold`}>List of my products:</h2>
                    </div>

                    <div className={`col d-flex justify-content-end`}>
                        <button className={`btn btn-outline-dark`} onClick={this.openModal}>Dodaj novi proizvod</button>
                    </div>

                </div>
                <div className={`mt-4 row`}>
                    {this.renderProductsCards()}
                </div>

                <ProductDetailModal
                    showModal={this.state.showModal}
                    productModel= {this.state.productModel}
                    productModelValidation={this.state.productModelValidation}
                    closeModal={this.closeModal}
                    saveButton={this.createProduct}
                />


            </div>
        );
    }
}

export default MyProducts;
