import { createContext} from "react";
import {ModalData} from "../components/ModalContainer";

type ModalContextData = {
    modal: ModalData | null
    setModal: (modals: ModalData | null ) => void
}

const ModalContext = createContext<ModalContextData>({
    modal: null,
    setModal: (): void => {return;},
});

export default ModalContext;