package apisix

//func resourceSsl() *schema.Resource {
//	return &schema.Resource{
//		Create: resourceSslCreate,
//		Read:   resourceSslRead,
//		Update: resourceSslUpdate,
//		Delete: resourceSslDelete,
//		Importer: &schema.ResourceImporter{
//			State: schema.ImportStatePassthrough,
//		},
//
//		Schema: map[string]*schema.Schema{
//			"certificate": {
//				Type:     schema.TypeString,
//				Required: true,
//			},
//			"private_key": {
//				Type:      schema.TypeString,
//				Required:  true,
//				Sensitive: true,
//			},
//		},
//	}
//}
//
//func resourceSslCreate(d *schema.ResourceData, m interface{}) error {
//	ssl, err := createSslObject(d)
//
//	if err != nil {
//		return err
//	}
//
//	result, err := m.(ApiClient).CreateSsl(ssl)
//	if err != nil {
//		return err
//	}
//
//	d.SetId(result.ID)
//	if err := d.Set("certificate", ssl.SslCertificate); err != nil {
//		return err
//	}
//	if err := d.Set("private_key", ssl.PrivateKey); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func resourceSslRead(d *schema.ResourceData, m interface{}) error {
//	id := d.Id()
//	ssl, err := m.(ApiClient).GetSsl(id)
//
//	if err != nil {
//		return err
//	}
//
//	if ssl.SslCertificate == "" {
//		d.SetId("")
//		return nil
//	}
//
//	return updateSslResource(d, ssl)
//}
//
//func resourceSslUpdate(d *schema.ResourceData, m interface{}) error {
//	id := d.Id()
//	ssl, err := createSslObject(d)
//
//	if err != nil {
//		return err
//	}
//
//	_, err = m.(ApiClient).UpdateSsl(id, ssl)
//	return err
//}
//
//func resourceSslDelete(d *schema.ResourceData, m interface{}) error {
//	id := d.Id()
//
//	err := m.(ApiClient).DeleteSsl(id)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func createSslObject(d *schema.ResourceData) (SSL, error) {
//	snis, err := ParseCert(d.Get("certificate").(string), d.Get("private_key").(string))
//
//	if err != nil {
//		return SSL{}, err
//	}
//
//	return SSL{
//		SslCertificate: d.Get("certificate").(string),
//		PrivateKey:  d.Get("private_key").(string),
//		SNIS:        snis,
//	}, nil
//}
//
//func updateSslResource(d *schema.ResourceData, data SSL) error {
//	return d.Set("certificate", data.SslCertificate)
//}
//
