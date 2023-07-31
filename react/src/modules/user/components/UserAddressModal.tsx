import React from "react";
import {useTranslation} from "react-i18next";
import {UserAddress} from "../../../__generated__/graphql";
import Button from "../../../common/components-library/Button";
import UserAddressForm from "./UserAddressForm";
import {Modal} from "../../../common/components/shelly-ui";
import {ModalHandler} from "../../../common/components/shelly-ui/Modal/useModal";

type UserAddressModalProps = {
    address?: UserAddress
    modal: ModalHandler
}

const UserAddressModal: React.FC<UserAddressModalProps> = (props) => {
    const {t} = useTranslation();

    return <Modal modal={props.modal}>
        <Modal.Title>
            {t("user_address.new_address_modal_title")}
        </Modal.Title>
        <UserAddressForm onSuccess={() => props.modal.close()} address={props.address}>
            <Modal.Actions>
                <Button onClick={() => props.modal.close()} htmlType="button">
                    {t("user_address.close_modal")}
                </Button>
                <Button type="primary" htmlType="submit">
                    {t("user_address.submit_form")}
                </Button>
            </Modal.Actions>
        </UserAddressForm>
    </Modal>;
};

export default UserAddressModal;