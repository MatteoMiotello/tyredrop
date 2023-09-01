import React, {useEffect, useState} from "react";
import {Button, Input, Textarea, useForm} from "./shelly-ui";
import Form from "./shelly-ui/Form";
import {useAuth} from "../../modules/auth/hooks/useAuth";
import {isRequired} from "./shelly-ui/Input";
import {createBackendClient} from "../backend/backend-client";
import {useToast} from "../../store/toast";

const ContactForm: React.FC = () => {
    const auth = useAuth();
    const form = useForm();
    const toastr = useToast();
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        form.setFormValues({
            email: auth.user?.user?.email
        });
    }, [auth.user]);

    const submit = (data: any) => {
        setLoading(true);
        return createBackendClient()
            .supportEmail({
                email: data.email,
                phone: data.phone,
                name: data.name,
                message: data.message
            })
            .then(() => toastr.success("Richiesta inviata con successo, verrai conttato presto da uno dei nostri operatori"))
            .catch(() => toastr.error("Non è stato possibile inviare la richiesta, riprova tra poco"))
            .finally(() => setLoading(false));
    };

    return <Form form={form} saveForm={submit}>
        <Input.FormControl>
            <Input type="tel" placeholder="Email" className="text-base-content my-1"
                   {...form.registerInput({
                name: 'email',
                validators: [isRequired( "L'email è richiesta" )],
            })}/>
        </Input.FormControl>
        <Input.FormControl>
            <Input placeholder="Telefono"
                   type="tel"
                   className="text-base-content my-1"
                   {...form.registerInput({name: 'phone', validators: [isRequired("Il telefono è richiesto")]})}
            />
        </Input.FormControl>
        <Input.FormControl>
            <Input placeholder="Nome"
                   type="tel"
                   className="text-base-content my-1"
                   {...form.registerInput({name: 'name', validators: [isRequired("Il nome è richiesto")]})}
            />
        </Input.FormControl>
        <Input.FormControl>
            <Textarea placeholder="Messaggio"
                      className="text-base-content my-1"
                      {...form.registerInput({name: 'message', validators: [isRequired("Il messaggio è richiesto")]})}
            />
        </Input.FormControl>
        <Form.FormButtons>
            <Button buttonType="primary" wide loading={loading}>
                Invia richiesta
            </Button>
        </Form.FormButtons>
    </Form>;
};

export default ContactForm;