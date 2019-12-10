import sys

sys.path.append(sys.path[0] + r'/../')

from prand import PRand
from hash import Pyrin

rand = PRand("test")

print(rand.rand())
print(rand.rand())
print(rand.rand())
print(rand.rand())

pyrin = Pyrin()
print(pyrin.to_hex(pyrin.hash("test")))
print(pyrin.to_hex(pyrin.hash("Lorem ipsum dolor sit amet, nunc leo consectetuer elit velit, id sapien, egestas consectetuer purus in, vel ipsum curabitur lorem amet. Enim non massa, a nulla et cras erat egestas. Tellus nec ipsum maecenas placerat, in curae lacinia. A donec. Duis ut dolor turpis eget suspendisse, lacus diam ante aliquam dolor posuere, dolor sed lacinia consequat augue condimentum, sollicitudin wisi tristique lectus vel. Pellentesque commodo. Non dui tellus nunc sed aliquam. Amet turpis tincidunt sapien vel duis. Tellus vel nam ipsum pulvinar, etiam pede tellus nam donec, mollis elementum. Sapien ac, per nec non magna risus, sed at donec sit fusce. Sed hendrerit vestibulum sed venenatis torquent, sed tellus scelerisque aliquam. Ut vel lectus, sed quis non sit justo, vitae quisque, non placerat metus, donec sed sagittis velit magna")))
