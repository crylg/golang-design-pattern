Golang Design Patterns
====

# 简介
设计模式（Design pattern）是一套被反复使用、多数人知晓的、经过分类编目的、代码设计经验的总结。使用设计模式是为了可重用代码、让代码更容易被他人理解、保证代码可靠性。 毫无疑问，设计模式于己于他人于系统都是多赢的，设计模式使代码编制真正工程化，设计模式是软件工程的基石，如同大厦的一块块砖石一样。

# 设计模式分类
<table>
  <tbody>
    <tr>
      <th width="40%">模式 &amp; 描述</th>
      <th>包括</th>
    </tr>
    <tr>
      <td><b>创建型模式</b><br>这些设计模式提供了一种在创建对象的同时隐藏创建逻辑的方式，而不是使用新的运算符直接实例化对象。这使得程序在判断针对某个给定实例需要创建哪些对象时更加灵活。
      </td>
      <td>
        <ul>
          <li><a href="creational/factory-method">工厂方法模式（Factory Pattern）</a></li>
          <li><a href="creational/abstract-factory">抽象工厂模式（Abstract Factory Pattern）</a></li>
          <li><a href="creational/singleton">单例模式（Singleton Pattern）</a></li>
          <li><a href="creational/builder">建造者模式（Builder Pattern）</a></li>
          <li><a href="creational/prototype">原型模式（Prototype Pattern）</a></li>
        </ul>
      </td>
    </tr>
    <tr>
      <td><b>结构型模式</b><br>这些设计模式关注类和对象的组合。继承的概念被用来组合接口和定义组合对象获得新功能的方式。</td>
      <td>
        <ul>
          <li><a href="adapter_pattern/adapter.md">适配器模式（Adapter Pattern）</a></li>
          <li><a href="bridge.md">桥接模式（Bridge Pattern）</a></li>
          <li><a href="filter.md">过滤器模式（Filter、Criteria Pattern）</a></li>
          <li><a href="composite.md">组合模式（Composite Pattern）</a></li>
          <li><a href="decorator.md">装饰器模式（Decorator Pattern）</a></li>
          <li><a href="facade.md">外观模式（Facade Pattern）</a></li>
          <li><a href="flyweight.md">享元模式（Flyweight Pattern）</a></li>
          <li><a href="proxy.md">代理模式（Proxy Pattern）</a></li>
        </ul>
      </td>
    </tr>
    <tr>
      <td><b>行为型模式</b><br>这些设计模式特别关注对象之间的通信。</td><td>
      <ul>
        <li><a href="responsibility_pattern/responsibility_pattern.md">责任链模式（Chain of Responsibility Pattern）</a></li>
        <li><a href="command.md">命令模式（Command Pattern）</a></li>
        <li><a href="interpreter.md">解释器模式（Interpreter Pattern）</a></li>
        <li><a href="iterator.md">迭代器模式（Iterator Pattern）</a></li>
        <li><a href="mediator.md">中介者模式（Mediator Pattern）</a></li>
        <li><a href="memento.md">备忘录模式（Memento Pattern）</a></li>
        <li><a href="observer_pattern\Observer.md">观察者模式（Observer Pattern）</a></li>
        <li><a href="state.md">状态模式（State Pattern）</a></li>
        <li><a href="null-object.md">空对象模式（Null Object Pattern）</a></li>
        <li><a href="strategy.md">策略模式（Strategy Pattern）</a></li>
        <li><a href="template_pattern/Template_Pattern.md">模板模式（Template Pattern）</a></li>
        <li><a href="visitor.md">访问者模式（Visitor Pattern）</a></li>
      </ul>
      </td>
    </tr>
  </tbody>
</table>
