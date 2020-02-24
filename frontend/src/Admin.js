import React from 'react';
import Form from 'react-bootstrap/Form';
import Col from 'react-bootstrap/Col';
import Row from 'react-bootstrap/Row';
import Button from 'react-bootstrap/Button';
import Container from 'react-bootstrap/Container';
import Toast from './Toast.js'

import axios from 'axios';

/*
This is the react equivilent of the following HTML form
<form action="http://localhost:8000" method="post" enctype="multipart/form-data">
  <p><input type="text" name="id" value="">
  <p><input type="file" name="file">
  <p><button type="submit">Submit</button>
</form>
*/


class Admin extends React.Component {

    constructor(props) {
        super(props);
        this.state = {validated: false, id: "", buttonDisabled: false, toastShow: false, toastText: "asd"};
    
        this.validated = this.validated.bind(this);
        this.changeHandler = this.changeHandler.bind(this);
        this.handleSubmit = this.handleSubmit.bind(this);
    }

    validated(){
        console.log("validated", this.state.validated)
        return this.state.validated;
    }

    handleSubmit(event) {
        event.preventDefault();
        if(event.currentTarget.checkValidity() === false) {
            event.stopPropagation();
            return;
        };

        this.setState({buttonDisabled: true, toastShow: false})

        // create the data
        const data = new FormData()
        data.append('file', this.state.file);
        data.append('id', this.state.id);

        // upload the file
        axios.post(
            window.global.files_location, 
            data, 
            {'content-type': `multipart/form-data; boundary=${data._boundary}`})
        .then(res => {
            console.log(res);
            var toastText = "";
            if(res.status === 200) {
                toastText = "Uploaded file";
            } else {
                toastText = "Unable to upload file. Error:" +res.statusText;
            }

            this.setState({buttonDisabled: false, toastShow: true, toastText: toastText});
        }).catch(error => {
            console.log("Err" + error);
            this.setState({buttonDisabled: false, toastShow: true, toastText: "Unable to upload file. " + error});
        });
    }

    changeHandler(event) {
        if(event.target.name === "file"){
            this.setState({ [event.target.name]: event.target.files[0], toastShow: false});
            return
        }

        this.setState({ [event.target.name]: event.target.value, toastShow: false});
    }

    render() {
      return (
          <div>
          <h1 style={{marginBottom: "40px"}}>Admin</h1>
          <Container className="text-left">
            <Form noValidate validated={this.validated} onSubmit={this.handleSubmit}>
                  <Form.Group as={Row} controlId="productID">
                      <Form.Label column sm="2">Product ID:</Form.Label>
                      <Col sm="6">
                          <Form.Control type="text" name="id" placeholder="" required style={{width: "80px"}} value={this.state.id} onChange={this.changeHandler}/>
                          <Form.Text className="text-muted">Enter the product id to upload an image for</Form.Text>
                          <Form.Control.Feedback type="invalid">Please provide a product ID.</Form.Control.Feedback>
                      </Col>
                  <Col sm="4">
                      <Toast show={this.state.toastShow} message={this.state.toastText}/>
                  </Col>
                  </Form.Group>
                  <Form.Group as={Row}>
                      <Form.Label column sm="2">File:</Form.Label>
                      <Col sm="10">
                          <Form.Control type="file" name="file" placeholder="" required onChange={this.changeHandler}/>
                          <Form.Text className="text-muted">Image to associate with the product</Form.Text>
                          <Form.Control.Feedback type="invalid">Please select a file to upload.</Form.Control.Feedback>
                      </Col>
                  </Form.Group>
                  <Button  type="submit" disabled={this.state.buttonDisabled}>Submit form</Button>
            </Form>
          </Container>
          </div>
      )
    }
}

export default Admin;