import React, {useState} from "react";
import Panel from "../../common/components-library/Panel";
import {Button, Form, Input, useForm} from "../../common/components/shelly-ui";
import {isRequired} from "../../common/validation/validators";
import {Link, useNavigate, useParams} from "react-router-dom";
import {z} from "zod";
import {zodParser} from "../../common/validation/zod-parser";
import {useToast} from "../../store/toast";
import {createBackendClient} from "../../common/backend/backend-client";
import {Simulate} from "react-dom/test-utils";
import load = Simulate.load;
import Spinner from "../../common/components/Spinner";

const ChangePasswordPage: React.FC = () => {
    const form = useForm();
    const {token} = useParams<{ token: string }>();
    const toastr = useToast();
    const navigate = useNavigate();
    const [loading, setLoading] = useState(false);

    const submit = (data: any) => {
        if (!token) {
            form.handleFormError("La richiesta non può essere elaborata, si prega di riprovare");
            return false;
        }

        setLoading(true);
        return createBackendClient()
            .changePassword(token, data.password)
            .then(() => {
                toastr.success("Password modificata con successo");
                navigate('/auth/login');
            })
            .catch(() => toastr.error("Non è stato possibile modificare la password"))
            .finally(() => setLoading(false));
    };

    return <>
        <div className="flex flex-col justify-center items-center my-auto ">
            {loading && <Spinner></Spinner>}
            <Panel className="md:w-1/2 w-full">
                <Form form={form} saveForm={submit}>
                    <Input.FormControl>
                        <Input.Label>
                            Password
                        </Input.Label>
                        <Input type="password" {...form.registerInput({
                            name: "password", validators: [
                                isRequired("password"),
                                (value) => {
                                    const passwordSchema = z.string()
                                        .min(8, {message: "La password deve contenere almeno 8 caratteri"});

                                    return zodParser(passwordSchema, value);
                                }
                            ]
                        })}></Input>
                    </Input.FormControl>
                    <Input.FormControl>
                        <Input.Label>
                            Ripeti password
                        </Input.Label>
                        <Input type="password" {...form.registerInput({
                            name: "repeat_password",
                            validators: [(value) => {
                                if (value != form.state.formValues?.getFormValue("password")) {
                                    return "Le due password non corrispondono";
                                }

                                return null;
                            }]
                        })} />
                    </Input.FormControl>
                    <Form.FormButtons>
                        <Link to="/auth/login" className="btn"> Torna al login </Link>
                        <Button type="submit" buttonType="primary">
                            Salva
                        </Button>
                    </Form.FormButtons>
                </Form>
            </Panel>
        </div>
    </>;
};

export default ChangePasswordPage;