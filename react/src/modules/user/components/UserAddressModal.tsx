import React from "react";
import {useTranslation} from "react-i18next";
import {UserAddress} from "../../../__generated__/graphql";
import Button from "../../../common/components-library/Button";
import Modal from "../../../common/components-library/Modal";
import UserAddressForm from "./UserAddressForm";

type UserAddressModalProps = {
    closeModal: () => void
    address?: UserAddress | undefined
}

const UserAddressModal: React.FC<UserAddressModalProps> = (props) => {
    const {t} = useTranslation();

    return <Modal.Content>
        <Modal.Header>
            {t("user_address.new_address_modal_title")}
        </Modal.Header>
        <UserAddressForm onSuccess={() => props.closeModal()} address={props.address}>
            <Modal.Action>
                <Button onClick={() => props.closeModal()} htmlType="button">
                    {t("user_address.close_modal")}
                </Button>
                <Button type="primary" htmlType="submit">
                    {t("user_address.submit_form")}
                </Button>
            </Modal.Action>
        </UserAddressForm>
    </Modal.Content>;
};

export default UserAddressModal;