import { Form, Input, Modal, Select, Switch } from 'antd';
import React, { useEffect } from 'react';
import { Device } from 'shared/types/device';

interface DeviceModalProps {
    open: boolean;
    device: Device | null;
    onClose: () => void;
    onSave: (values: Device) => Promise<void>;
    confirmLoading: boolean;
}

export const DeviceModal: React.FC<DeviceModalProps> = ({
    open,
    device,
    onClose,
    onSave,
    confirmLoading
}) => {
    const [form] = Form.useForm<Device>();

    useEffect(() => {
        if (open) {
            if (device) {
                form.setFieldsValue(device);
            } else {
                form.resetFields();
                form.setFieldsValue({ is_active: true, platform_code: 'LINUX' });
            }
        }
    }, [open, device, form]);

    const handleOk = async () => {
        try {
            const values = await form.validateFields();
            await onSave(values);
        } catch (error) {
            // Валидация не прошла
        }
    };

    return (
        <Modal
            title={device ? 'Редактировать устройство' : 'Добавить новое устройство'}
            open={open}
            onOk={handleOk}
            onCancel={onClose}
            confirmLoading={confirmLoading}
            okText="Сохранить"
            cancelText="Отмена"
            destroyOnClose
        >
            <Form form={form} layout="vertical" name="deviceForm">
                <Form.Item
                    name="hostname"
                    label="Hostname"
                    rules={[{ required: true, message: 'Пожалуйста, введите hostname устройства' }]}
                >
                    <Input placeholder="например, workstation-01" />
                </Form.Item>

                <Form.Item
                    name="ip"
                    label="IP Адрес"
                    rules={[
                        { required: true, message: 'Введите IP адрес' },
                        {
                            pattern: /^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$/,
                            message: 'Некорректный формат IP'
                        }
                    ]}
                >
                    <Input placeholder="192.168.1.10" />
                </Form.Item>

                <Form.Item name="platform_code" label="Платформа" rules={[{ required: true }]}>
                    <Select>
                        <Select.Option value="LINUX">Linux OS</Select.Option>
                        <Select.Option value="WINDOWS">Windows OS</Select.Option>
                        <Select.Option value="MACOS">macOS</Select.Option>
                    </Select>
                </Form.Item>

                <Form.Item name="is_active" label="Активно" valuePropName="checked">
                    <Switch />
                </Form.Item>
            </Form>
        </Modal>
    );
};
