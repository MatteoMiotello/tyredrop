import {faPlus} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useEffect, useState} from "react";
import {useSelector} from "react-redux";
import Button from "../../../common/components-library/Button";
import Form from "../../../common/components-library/Form";
import Modal from "../../../common/components-library/Modal";
import {SelectComponent, SelectOption} from "../../../common/components-library/SelectComponent";
import useModal from "../../../hooks/useModal";
import UserAddressForm from "../../user/components/UserAddressForm";
import userSelector from "../../user/store/user-selector";

const UserAddressSelector: React.FC = () => {
    const userAddresses = useSelector(userSelector.addresses);
    const [addressOptions, setOptions] = useState<SelectOption[]>([]);

    useEffect(() => {
        if (!userAddresses || !userAddresses.length) {
            setOptions([
                {
                    title: "Nessun indirizzo selezionato",
                    value: 1,
                    disabled: true
                }
            ]);
            return;
        }

        const options = userAddresses.map((address) => {
            return {
                title: address.addressLine1,
                value: address.ID,
            };
        });


        setOptions(options);
    }, [userAddresses]);

    const {openModal, closeModal} = useModal(<Modal.Content>
            <Modal.Header>
                Inserisci un nuovo indirizzo
            </Modal.Header>
            <UserAddressForm>
                <Modal.Action>
                    <Button onClick={() => closeModal()}>
                        Annulla
                    </Button>
                    <Button type="primary" htmlType="submit">
                        Salva
                    </Button>
                </Modal.Action>
            </UserAddressForm>
        </Modal.Content>
    );

    return <div>
        <div>
            {
                (userAddresses && userAddresses.length) ?
                    <Form
                        onSubmit={}
                        form={}
                    >
                        <SelectComponent
                            className="mb-2"
                            options={addressOptions}
                            name="user_address"
                        /> : <span className="mb-4"> Nessun indirizzo trovato.. </span>
                    </Form>
            }
        </div>
        <a className="link link-secondary" onClick={openModal}>
            <FontAwesomeIcon icon={faPlus}/> Aggiungi un nuovo indirizzo
        </a>
    </div>;
};

export default UserAddressSelector;