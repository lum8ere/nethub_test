import { Form, Input, Modal, Select, Switch } from 'antd';
import React, { useEffect } from 'react';
import { Device } from 'shared/types/device';
import { Location } from 'shared/types/location';
import { Platform } from 'shared/types/platform';

interface DeviceModalProps {
    open: boolean;
    device: Device | null;
    locations: Location[];
    platforms: Platform[];
    onClose: () => void;
    onSave: (values: Device) => Promise<void>;
    confirmLoading: boolean;
}

export const DeviceModal: React.FC<DeviceModalProps> = ({
    open,
    device,
    locations,
    platforms,
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
                form.setFieldsValue({ is_active: true });
            }
        }
    }, [open, device, form]);

    const handleOk = async () => {
        const values = await form.validateFields();
        await onSave(values);
    };

    return (
        <Modal
            title={device ? 'Редактировать устройство' : 'Добавить устройство'}
            open={open}
            onOk={handleOk}
            onCancel={onClose}
            confirmLoading={confirmLoading}
            destroyOnHidden
        >
            <Form form={form} layout="vertical">
                <Form.Item name="hostname" label="Hostname" rules={[{ required: true }]}>
                    <Input />
                </Form.Item>
                <Form.Item name="ip" label="IP Адрес" rules={[{ required: true }]}>
                    <Input />
                </Form.Item>
                <Form.Item name="platform_code" label="Платформа" rules={[{ required: true }]}>
                    <Select placeholder="Выберите платформу">
                        {platforms.map((p) => (
                            <Select.Option key={p.code} value={p.code}>
                                {p.name}
                            </Select.Option>
                        ))}
                    </Select>
                </Form.Item>
                <Form.Item name="location" label="Локация">
                    <Select placeholder="Выберите локацию" allowClear>
                        {locations.map((l) => (
                            <Select.Option key={l.id} value={l.id}>
                                {l.name}
                            </Select.Option>
                        ))}
                    </Select>
                </Form.Item>

                <Form.Item name="is_active" label="Активно" valuePropName="checked">
                    <Switch />
                </Form.Item>
            </Form>
        </Modal>
    );
};
