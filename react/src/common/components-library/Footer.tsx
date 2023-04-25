import React from "react";

export type FooterColumn = {
    title?: string
    key: number | string
    links: FooterLink[]
}

export type FooterLink = {
    title: string
    key: number | string
    url?: string
}

export interface FooterProps {
    data: FooterColumn[]
}

const Footer: React.FC<FooterProps> = ( props ) => {
    return <footer className="footer p-10 bg-neutral text-neutral-content w-full">
        {
            props.data.map( (column: FooterColumn) => <div key={column.key}>
                <span className="footer-title">{ column.title }</span>
                    { column.links.map( ( link: FooterLink ) => <a key={link.key} className="link link-hover" href={link.url}>{link.title}</a> ) }
                 </div> )
        }
    </footer>;
};

export default Footer;