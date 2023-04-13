import React from "react";

export type FooterColumn = {
    title?: string
    links: FooterLink[]
}

export type FooterLink = {
    title: string
    url?: string
}

export interface FooterProps {
    data: FooterColumn[]
}

const Footer: React.FC<FooterProps> = ( props ) => {
    return <footer className="footer p-10 bg-neutral text-neutral-content">
        {
            props.data.map( (column: FooterColumn) => <div>
                <span className="footer-title">{ column.title }</span>
                    { column.links.map( ( link: FooterLink ) => <a className="link link-hover" href={link.url}>{link.title}</a> ) }
                 </div> )
        }
    </footer>
}

export default Footer