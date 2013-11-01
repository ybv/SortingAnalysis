import string
import cProfile


def main():
	alist =[]
	swapcnt=0
	f = open('10MB', 'r+b')
	fw = open('bigbpy', 'r+b')
	for line in f:
		alist.append(line)
	
	def bubbleSort(alist):
		swapcnt =0
		for passnum in range(len(alist)-1,0,-1):
			for i in range(passnum):
				if cmp(alist[i],alist[i+1])==1:
					temp = alist[i]
					alist[i] = alist[i+1]
					alist[i+1] = temp 
					swapcnt=swapcnt+1
                		
	
	bubbleSort(alist)
	for l in alist:
		fw.write(str(l))
	print "number of swaps",swapcnt

if __name__ == '__main__':
    main()
    
