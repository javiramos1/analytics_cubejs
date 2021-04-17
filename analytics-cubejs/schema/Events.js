cube(`Events`, {
  refreshKey: {
    every: `15 second`
  },
  sql: `SELECT * FROM analytics.events where event_type = 'click'`,
  
  joins: {
  
  },
  
  measures: {
    count: {
      type: `count`,
      drillMembers: [eventId, clickId]
    },
  },
  
  dimensions: {
    eventId : {
      sql: `event_id`,
      type: `string`,
      primaryKey: true,
    },

    pagePath: {
      sql: `page_path`,
      type: `string`
    },
    
    clickText: {
      sql: `click_text`,
      type: `string`
    },
    
    clickId: {
      sql: `click_id`,
      type: `string`
    },
    
    clickClass: {
      sql: `click_class`,
      type: `string`
    },

    userIp: {
      sql: `user_ip`,
      type: `string`
    },
        
    userAgent: {
      sql: `user_agent`,
      type: `string`
    },

    userOs: {
      sql: `user_os`,
      type: `string`
    },
        
    isMobile: {
      sql: `is_mobile`,
      type: `boolean`
    },

    userCity: {
      sql: `user_city`,
      type: `string`
    },
        
    userCountryCode: {
      sql: `user_country_code`,
      type: `string`
    },

    userCountryName: {
      sql: `user_country_name`,
      type: `string`
    },

    userCountryRegion: {
      sql: `user_country_region`,
      type: `string`
    },

    userContinent: {
      sql: `user_continent`,
      type: `string`
    },

    userTimezone: {
      sql: `user_timezone`,
      type: `string`
    },

    userLanguages: {
      sql: `user_languages`,
      type: `string`
    },

    userInEu: {
      sql: `user_in_eu`,
      type: `boolean`
    },

    eventTime: {
      sql: `event_time`,
      type: `time`
    }
  },

  segments: {
    IsUserInEu: {
      sql:`${CUBE}.user_in_eu = true`
    },
    IsUserNotInEu: {
      sql:`${CUBE}.user_in_eu = false`
    },
    IsUserInMobile: {
      sql:`${CUBE}.is_mobile = true`
    },
    IsUserNotInMobile: {
      sql:`${CUBE}.is_mobile = false`
    }
  },

  dataSource: `default`
});



cube(`PageViews`, {
  refreshKey: {
    every: `15 second`
  },
  sql: `SELECT * FROM analytics.events where event_type = 'view'`,
  
  joins: {

  },
  
  measures: {
    count: {
      type: `count`,
      drillMembers: [eventId]
    }
  },
  
  dimensions: {
    eventId : {
      sql: `event_id`,
      type: `string`,
      primaryKey: true
    },

    pagePath: {
      sql: `page_path`,
      type: `string`
    },

    userIp: {
      sql: `user_ip`,
      type: `string`
    },
        
    userAgent: {
      sql: `user_agent`,
      type: `string`
    },

    userOs: {
      sql: `user_os`,
      type: `string`
    },
        
    isMobile: {
      sql: `is_mobile`,
      type: `boolean`
    },

    userCity: {
      sql: `user_city`,
      type: `string`
    },
        
    userCountryCode: {
      sql: `user_country_code`,
      type: `string`
    },

    userCountryName: {
      sql: `user_country_name`,
      type: `string`
    },

    userCountryRegion: {
      sql: `user_country_region`,
      type: `string`
    },

    userContinent: {
      sql: `user_continent`,
      type: `string`
    },

    userTimezone: {
      sql: `user_timezone`,
      type: `string`
    },

    userLanguages: {
      sql: `user_languages`,
      type: `string`
    },

    userInEu: {
      sql: `user_in_eu`,
      type: `boolean`
    },
    
    eventTime: {
      sql: `event_time`,
      type: `time`
    }
  },

  segments: {
    IsUserInEu: {
      sql:`${CUBE}.user_in_eu = true`
    },
    IsUserNotInEu: {
      sql:`${CUBE}.user_in_eu = false`
    },
    IsUserInMobile: {
      sql:`${CUBE}.is_mobile = true`
    },
    IsUserNotInMobile: {
      sql:`${CUBE}.is_mobile = false`
    }
  },


  
  dataSource: `default`
});
