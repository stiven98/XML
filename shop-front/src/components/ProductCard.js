import React from "react";
import {Button, Card} from "react-bootstrap";
import {withRouter} from "react-router-dom";
import ImagesApi from "../api/ImagesApi";

class ProductCard extends React.Component {


    onProductPage = () => {
        this.props.history.push('/product/' + this.props.product.id);
    }

    render() {
        return (
            <div className={`col-4 mt-2 mb-2`}>
                <Card style={{width: '18rem'}}>
                    <Card.Img variant="top" src={ImagesApi + this.props.product.picturePath}/>
                    <Card.Body>
                        <Card.Title>
                            <div className={`row`}>
                                <div className={`col d-flex justify-content-start`}>
                                    {this.props.product.name}
                                </div>
                                <div className={`col d-flex justify-content-end text-success`}>
                                    {this.props.product.price} $

                                </div>
                            </div>

                        </Card.Title>

                        <div className={`d-flex justify-content-center`}>
                            <Button variant="primary" onClick={this.onProductPage}>Pogledaj opis</Button>

                        </div>
                    </Card.Body>
                </Card>
            </div>


        );
    }
}

export default withRouter(ProductCard);
