import React from "react";
import {Button, Modal} from "react-bootstrap";
import ProductModelValidation from "../model/ProductModelValidation";
import axios from "axios";
import ImagesApi from "../api/ImagesApi";

class ProductDetailModal extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            productModel: this.props.productModel,
            productModelValidation: this.props.productModelValidation,
        };
    }



    changeName = (event) => {
        this.setState({productModel: {...(this.state.productModel), name: event.target.value}});
        this.setState({productModelValidation: new ProductModelValidation()})
    }

    changePrice = (event) => {
        this.setState({productModel: {...(this.state.productModel), price: event.target.value}});
        this.setState({productModelValidation: new ProductModelValidation()})
    }

    changeQuantity = (event) => {
        this.setState({productModel: {...(this.state.productModel), quantity: event.target.value}});
        this.setState({productModelValidation: new ProductModelValidation()})
    }

    changePhoto = (event) => {
        event.preventDefault();
        this.setState({productModelValidation: new ProductModelValidation()})
        const formData = new FormData();
        formData.append('files', event.target.files[0]);
        axios.post(ImagesApi + 'upload', formData).then(response => {
            if (response.data !== '') {
                this.setState({productModel: {...this.state.productModel, picturePath: response.data}});
            }
        });
    }

    renderPhoto = () => {
        if (this.state.productModel.picturePath !== '') {
            return (<img alt={`Product`} src={ImagesApi + this.state.productModel.picturePath} style={{ width: '100%', height:'100%' }}/>);
        }
    }


    isValidName = (name) => {
        if (name.length > 0) {
            this.setState({productModelValidation : {...this.state.productModelValidation, validName: 'is-valid'}});
            return true;
        } else {
            this.setState({productModelValidation : {...this.state.productModelValidation, validName: 'is-invalid'}});
            return false;
        }
    }

    isValidPrice = (price) => {
        if (price !== '') {
            this.setState({productModelValidation : {...this.state.productModelValidation, validPrice: 'is-valid'}});
            return true;
        } else {
            this.setState({productModelValidation : {...this.state.productModelValidation, validPrice: 'is-invalid'}});
            return false;
        }
    }

    isValidQuantity = (quantity) => {
        if (quantity !== '') {
            this.setState({productModelValidation : {...this.state.productModelValidation, validQuantity: 'is-valid'}});
            return true;
        } else {
            this.setState({productModelValidation : {...this.state.productModelValidation, validQuantity: 'is-invalid'}});
            return false;
        }
    }

    isValidPicturePath = (path) => {
        if (path.length > 0) {
            this.setState({productModelValidation : {...this.state.productModelValidation, validPicturePath: 'is-valid'}});
            return true;
        } else {
            this.setState({productModelValidation : {...this.state.productModelValidation, validPicturePath: 'is-invalid'}});
            return false;
        }
    }

    save = async (event) => {
        event.preventDefault();
        if (await this.isFormValid()) {
            this.setState({productModel: {...(this.state.productModel), userID: localStorage.getItem('id')}});
            this.setState({productModel: {...(this.state.productModel), price: parseFloat(this.state.productModel.price)}});
            this.setState({productModel: {...(this.state.productModel), quantity: parseInt(this.state.productModel.quantity)}});
            console.log(this.state.productModel);
            this.props.saveButton(this.state.productModel);
        }
    }

    closeModal = () => {
        this.props.closeModal();
        this.setState({
            productModel: this.props.productModel,
            productModelValidation: this.props.productModelValidation,
        });
    }

    isFormValid = async () => {
        const validName = await this.isValidName(this.state.productModel.name);
        const validQuality = await this.isValidQuantity(this.state.productModel.quantity);
        const validPrice = await this.isValidPrice(this.state.productModel.price);
        const validPath = await this.isValidPicturePath(this.state.productModel.picturePath);
        return validName && validQuality && validPrice && validPath;
    }

    render() {
        return (
            <Modal show={this.props.showModal}>
                <Modal.Header>
                    <Modal.Title>Novi proizvod:</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <form className="form-control-feedback">
                        <div className={`row mt-3`}>
                            {this.renderPhoto()}
                        </div>

                        <div className={`row mt-3`}>
                            <div className={`col-12 d-flex justify-content-center`}>
                                <label htmlFor={`upload`} className={`mt-2 mr-2`}>Ucitajte sliku za
                                    proizvod:</label>
                                <input id="upload" className={`form-control w-50 ` + this.state.productModelValidation.validPicturePath } type="file" name="file" onChange={this.changePhoto}/>
                            </div>
                        </div>
                        <div className={`row d-flex justify-content-center mt-3`}>
                            <div className={`w-50`}>
                                <label htmlFor="name" className="text-dark">Naziv proizvoda:</label>
                                <input type="text" id="name" className={`form-control ` + this.state.productModelValidation.validName } value={this.state.productModel.name} onChange={this.changeName}/>
                                <div className="invalid-feedback">
                                    Input product!
                                </div>
                            </div>
                        </div>
                        <div className={`row d-flex justify-content-center mt-3`}>
                            <div className={`w-50`}>
                                <label htmlFor="price" className="text-dark">Cena:</label>
                                <input type="number" id="price" className={`form-control ` + this.state.productModelValidation.validPrice } value={this.state.productModel.price} onInput={this.changePrice}/>
                                <div className="invalid-feedback">
                                    Input price!
                                </div>
                            </div>
                        </div>

                        <div className={`row d-flex justify-content-center mt-3`}>
                            <div className={`w-50`}>
                                <label htmlFor="quantity" className="text-dark">Kolicina:</label>
                                <input type="number" id="quantity" className={`form-control ` + this.state.productModelValidation.validQuantity }  value={this.state.productModel.quantity} onInput={this.changeQuantity}/>
                                <div className="invalid-feedback">
                                    Input quantity!
                                </div>
                            </div>
                        </div>
                    </form>
                </Modal.Body>

                <Modal.Footer>
                    <Button variant="secondary" onClick={this.closeModal}>Zatvori</Button>
                    <Button variant="primary" onClick={this.save}>Sacuvaj</Button>
                </Modal.Footer>
            </Modal>
        );
    }
}

export default ProductDetailModal;
