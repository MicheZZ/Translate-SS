def dec_to_base(N, base):
    if not hasattr(dec_to_base, 'table'):
        dec_to_base.table = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ'
    x,y = divmod(N, base)
    return dec_to_base(x, base) + dec_to_base.table[y] if x else dec_to_base.table[y]


key3 = None
while key3 != True:
    try:
        first = int(input("Из какой системы счисления надо будет переводить? - "))
        if first >= 2 and first <= 36:
            key3 = True
        else:
            print("Введи целое число от 2 до 36")
    except:
        print("Введи целое число от 2 до 36")

key2 = None
while key2 != True:
    try:
        second = int(input("В какую систему счисления надо будет переводить? - "))
        if second >= 2 and second <= 36:
            key2 = True
        else:
            print("Введи целое число от 2 до 36")
    except:
        print("Введи целое число от 2 до 36")

key = None
while key != True:
    x = input("Какое число надо будет переводить? - ")
    try:
        g = int(x, first)
        key = True
    except ValueError:
        print("Введи любое целое число. Примечание: Цифры в числе не должны быть меньше системы счисления, из которой переводишь")


print(dec_to_base(g, second))


input()