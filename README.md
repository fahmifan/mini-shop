# Event Bus
- the purpose of this project is to create a decouple monilithic service
- the sample app is a an ecommerce app, with following features:
    - order
    - payment
    - listing product

## Designs
We will create the design of data flow and how each part of service communicate using
event bus.
```

+-------------+
|create order +----------+
+-------------+          |
                    +----v----+
                    |         |
                    |   BUS   |
                    |         |
                    +----^----+
                         |
                         |
+--------------+         |
|create payment+---------+
+--------------+         |
                         |
                         |
+--------------+         |
|create notif  +---------+
+--------------+

```

### Folder structure
```
.
├── README.md
├── eventbus
│   ├── bus.go
│   └── topic.go
├── eventhandler
│   └── handler.go
├── go.mod
├── go.sum
├── main.go
├── model
│   └── model.go
└── service
    ├── order_service.go
    ├── payment_service.go
    ├── product_service.go
    └── service.go
```

## Notes
- [x] add http web service