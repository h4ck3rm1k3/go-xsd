
test: go-xsd
	#./go-xsd --gofmt=False --goinst=False  --uri file:///mnt/data/home/mdupont/experiments/odoo/openerp/addons/base/rng/view.xsd
	#./go-xsd --gofmt=False --goinst=False --uri file:///mnt/data/home/mdupont/experiments/odoo/openerp/import_xml.xsd
	./go-xsd --gofmt=False --goinst=False --uri file:///mnt/data/home/mdupont/experiments/odoo/openerp/test.xsd

encoding2/xml.o : encoding2/xml/marshal.go encoding2/xml/read.go encoding2/xml/typeinfo.go encoding2/xml/xml.go
	gccgo-5 -g -c -o encoding2/xml.o $^

go-xsd : types.o  ufs.o  ugo.o  uslice.o  ustr.o  xsd.o xsd-makepkg/main.go unet.o encoding2/xml.o
	gccgo-5 -o go-xsd  -g $^ 

xsd.o : ./elem.go ./elemmakepkg.go ./elemparents.go ./hasattr.go ./haselem.go ./hasmakepkg.go ./hasparents.go ./makepkg.go ./xsd-schema.go
	gccgo-5 -c -g -o xsd.o $^

types.o : ./types/xsdtypes.go
	gccgo-5 -c -g -o types.o $^
