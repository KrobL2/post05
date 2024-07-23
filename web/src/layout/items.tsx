import React from "react";
import {
  AuditOutlined,
  BankOutlined,
} from "@ant-design/icons";

export const items = [
  {
    key: "sub1",
    icon: <BankOutlined />,
    label: "Дефекты",
    children: [
      {
        key: 1,
        label: "Квартиры",
        test: "232323",
        link: "/?path=/story/example-introduction--page",
      },
      {
        key: 2,
        test: "232323",
        label: "Инженерные сети",
        link: "link2",
      },
    ],
  },
  {
    key: "sub2",
    icon: <AuditOutlined />,
    label: "subnav 2",
    children: [
      {
        key: 5,
        label: "Ntcn",
      },
      {
        key: 6,
        label: "Tfdfdjf",
      },
      {
        key: 7,
        label: "ereere",
      },
    ],
  },
  {
    key: "sub3",
    label: "subnav 3",
    children: [
      {
        key: 9,
        label: "Ntcn",
      },
      {
        key: 10,
        label: "Tfdfdjf",
      },
      {
        key: 11,
        label: "ereere",
      },
    ],
  },
];
