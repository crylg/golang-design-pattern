抽象工厂模式
====
# 简介
**抽象工厂模式**是围绕一个超级工厂创建其他工厂。该超级工厂又称为其他工厂的工厂。这种类型的设计模式属于创建型模式，它提供了一种创建对象的最佳方式。

# 意图
定义一个用于创建对象的接口，让子类决定实例化哪一个类，工厂方法使一个类的实例化延迟到其子类。

# 抽象工厂模式有四个要素
+ **产品接口(Driver)** 定义产品的规范，所有的产品实现都必须遵循产品接口定义的规范。产品接口是调用者最为关心的，产品接口定义的优劣直接决定了调用者代码的稳定性。

+ **产品实现(Mysql/Oracle/Sqlite)** 产品接口的具体实现类，不同的产品也需要不同的产品实现类，产品实现类与功能创建类相对应。
  
+ **工厂接口(AbstractFactory)** 工厂方法模式的核心，与调用者直接交互用来提供产品的创建。

+ **工厂实现(MysqlFactory/OracleFactory/SqliteFactory)** 决定如何实例化产品，是实现扩展的途径，需要有多少种产品，就有多少个具体的工厂实现类，每个工厂实现类负责创建一种产品。


# 优点
1. 每增加一个产品就增加一个对应的工厂来创建它，这样整个工厂和产品体系都没有什么变化，而只是扩展的变化，这就完全符合开放-封闭的原则了
2. 严格遵循面向对象类的设计原则，比如单一职能原则、开-闭原则、依赖倒置原则、迪米特原则。
3. 业务实现解耦: 抽象工厂是静态工厂方法的进一步抽象与推广，由于使用了多态性，抽象工厂模式保持了静态工厂方法的优点同时又克服了它的缺点，不过抽象工厂模式自己的缺点是每加一个产品都需要增加一个工厂类，增加了大量的开发工作量。

# 缺点
1. 产品族扩展非常困难，要增加一个系列的某一产品，既要在抽象的 Creator 里加代码，又要在具体的里面加代码。