import React from "react";
import Footer, {FooterColumn} from "../components-library/Footer";
import {useTranslation} from "react-i18next";

const CustomFooter: React.FC = () => {
    const {t} = useTranslation();

    const footerData: FooterColumn[] = [
        {
            key:1,
            links: [
                {
                    title: t('footer.who'),
                    key: 1,
                },
                {
                    title: t('footer.newsletter'),
                    key: 2,
                },
                {
                    title: t('footer.assistance'),
                    key: 3
                },
                {
                    title: t('footer.faq'),
                    key: 4
                }
            ]
        }
    ];

    return <Footer data={footerData}></Footer>;
};
export default CustomFooter;