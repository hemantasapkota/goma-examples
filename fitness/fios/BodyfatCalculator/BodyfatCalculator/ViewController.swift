//
//  ViewController.swift
//  KetoCalculator
//
//  Created by Hemanta Sapkota on 16/10/16.
//  Copyright © 2016 Hemanta Sapkota. All rights reserved.
//

import UIKit
import SnapKit
import SwiftListView

class ViewController: UIViewController {
    
    override func viewDidLoad() {
        super.viewDidLoad()
        // Do any additional setup after loading the view, typically from a nib.
        
        self.view = BodyFatView()
    }

    override func didReceiveMemoryWarning() {
        super.didReceiveMemoryWarning()
        // Dispose of any resources that can be recreated.
    }

}

class BodyFatView : UIView {
    
    var metricBtn: GroupFilterButton!
    
    var genderBtn: GroupFilterButton!
    
    var heightBtn: GroupFilterButton!
    
    var neckBtn: GroupFilterButton!
    
    var waistBtn: GroupFilterButton!
    
    var hipBtn: GroupFilterButton!
    
    init() {
        super.init(frame: UIScreen.main.bounds)
        
        self.backgroundColor = UIColor.white
        
        // Metric
        metricBtn = GroupFilterButton()
        metricBtn.selectedText = "Choose a metric"
        metricBtn.onSelection = { selected in
            self.makeListView(title: "Choose a metric", highlighted: selected, items: ["Inches", "Centimeters"], filterBtn: self.metricBtn)
        }
        
        addSubview(metricBtn)
        metricBtn.snp.makeConstraints { (make) in
            make.top.equalTo(50)
            make.centerX.equalTo(self.snp.centerX)
            
            make.width.equalTo(self.snp.width).offset(-20)
            make.height.equalTo(40)
        }
        
        // Gender
        genderBtn = GroupFilterButton()
        genderBtn.selectedText = "Select a gender"
        genderBtn.onSelection = { selected in
            self.makeListView(title: "Select a gender", highlighted: selected, items: ["Male", "Female"], filterBtn: self.genderBtn)
        }
        addSubview(genderBtn)
        genderBtn.snp.makeConstraints { (make) in
            make.top.greaterThanOrEqualTo(metricBtn.snp.bottom).offset(5)
            make.centerX.equalTo(self.snp.centerX)
            
            make.width.equalTo(self.snp.width).offset(-20)
            make.height.equalTo(40)
        }
        
        // Height
        heightBtn = GroupFilterButton()
        heightBtn.selectedText = "Enter height"
        heightBtn.onSelection = { selected in
            
//            let alertController = UIAlertController(title: "Title", message: "Message", preferredStyle: .alert)
//            self.presentViewController(alertController, animated: true)
            
        }
        
        addSubview(heightBtn)
        heightBtn.snp.makeConstraints { (make) in
            make.top.greaterThanOrEqualTo(genderBtn.snp.bottom).offset(5)
            make.centerX.equalTo(self.snp.centerX)
            
            make.width.equalTo(self.snp.width).offset(-20)
            make.height.equalTo(40)
        }
        
        // Neck mesaurement
        neckBtn = GroupFilterButton()
        neckBtn.selectedText = "Enter neck measurement"
        addSubview(neckBtn)
        neckBtn.snp.makeConstraints { (make) in
            make.top.greaterThanOrEqualTo(heightBtn.snp.bottom).offset(5)
            make.centerX.equalTo(self.snp.centerX)
            
            make.width.equalTo(self.snp.width).offset(-20)
            make.height.equalTo(40)
        }
        
        // Wasit measurement
        waistBtn = GroupFilterButton()
        waistBtn.selectedText = "Enter waist measurement"
        addSubview(waistBtn)
        waistBtn.snp.makeConstraints { (make) in
            make.top.greaterThanOrEqualTo(neckBtn.snp.bottom).offset(5)
            make.centerX.equalTo(self.snp.centerX)
            
            make.width.equalTo(self.snp.width).offset(-20)
            make.height.equalTo(40)
        }
        
        // Hip measurement
        hipBtn = GroupFilterButton()
        hipBtn.selectedText = "Enter hip measurement"
        addSubview(hipBtn)
        hipBtn.snp.makeConstraints { (make) in
            make.top.greaterThanOrEqualTo(waistBtn.snp.bottom).offset(5)
            make.centerX.equalTo(self.snp.centerX)
            
            make.width.equalTo(self.snp.width).offset(-20)
            make.height.equalTo(40)
        }       
    }
    
    func makeListView(title:String, highlighted: String, items: [String], filterBtn: GroupFilterButton) {
        let listView = BasicListView(viewTitle: title, highlighted: highlighted)
        listView.Items = items
        listView.onSelection = { selected in
            filterBtn.selectedText = selected
        }
        listView.show()
    }
    
    required init?(coder aDecoder: NSCoder) {
        fatalError("init(coder:) has not been implemented")
    }
    
}

