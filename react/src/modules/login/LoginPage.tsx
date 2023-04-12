import React from "react";
import {Button, Form, Input, Layout, Space} from "antd";
import {Content, Footer} from "antd/es/layout/layout";

export const LoginPage: React.FC = () => {
    return <Space className="m-0" direction="vertical" style={{width: '100%'}}>
        <Layout className="h-screen">
            <Content className="h-full p-24">
                <Form
                    name="login"
                    layout="vertical"
                    size="large"
                    className={"flex flex-col items-center"}
                >
                    <Form.Item
                        name="username"
                        className="mb-3 w-md-1/2 my-auto"
                        rules={[{required: true, message:"Username is required"}]}
                    >
                        <Input placeholder="Username"/>
                    </Form.Item>
                    <Form.Item
                        name="password"
                        rules={[{required: true, message:"Password is required"}]}
                        rootClassName="mb-3 w-md-1/2 my-auto"
                    >
                        <Input.Password placeholder="Password"/>
                    </Form.Item>
                    <Form.Item>
                        <Button
                            htmlType="submit"
                            type="primary"
                            className="mx-2"
                        >
                            Login
                        </Button>
                    </Form.Item>
                </Form>
            </Content>
            <Footer>Footer</Footer>
        </Layout>
    </Space>
}