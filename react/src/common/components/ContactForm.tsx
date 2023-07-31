import React, {useEffect} from "react";
import {Button, Input, Textarea, useForm} from "./shelly-ui";
import Form from "./shelly-ui/Form";
import {useAuth} from "../../modules/auth/hooks/useAuth";

const ContactForm: React.FC = () => {
    const auth = useAuth();
    const form = useForm();
    useEffect( () => {
        form.setFormValues( {
            email: auth.user?.user?.email
        } );
    }, [auth.user] );

    return <Form form={form} saveForm={ (data) => false }>
        <Input.FormControl>
            <Input type="tel" placeholder="Email" className="text-base-content my-1" {...form.registerInput({name: 'email', disable: Boolean( auth.user?.user )} )}/>
        </Input.FormControl>
        <Input.FormControl>
            <Input placeholder="Telefono" type="tel" className="text-base-content my-1" name="phone"/>
        </Input.FormControl>
        <Input.FormControl>
            <Textarea placeholder="Messaggio" className="text-base-content my-1" name="message"/>
        </Input.FormControl>
        <Form.FormButtons>
            <Button buttonType="primary" wide>
                Invia richiesta
            </Button>
        </Form.FormButtons>
    </Form>;
};

export default ContactForm;